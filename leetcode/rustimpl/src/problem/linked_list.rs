use std::{cell::RefCell, rc::Rc};

type Link<T> = Option<Rc<RefCell<Node<T>>>>;

#[derive(Debug)]
pub struct Node<T> {
    val: T,
    prev: Link<T>,
    next: Link<T>,
}

impl<T> Node<T> {
    pub fn new(val: T) -> Rc<RefCell<Node<T>>> {
        Rc::new(RefCell::new(Self {
            val,
            prev: None,
            next: None,
        }))
    }
}

#[derive(Debug)]
pub struct LinkedList<T> {
    length: usize,
    head: Link<T>,
    tail: Link<T>,
}

impl<T> LinkedList<T> {
    pub fn new() -> Self {
        Self {
            length: 0,
            head: None,
            tail: None,
        }
    }

    pub fn push_front(&mut self, val: T) {
        let new_head = Node::new(val);
        match self.head.take() {
            Some(old_head) => {
                old_head.borrow_mut().prev = Some(new_head.clone());
                new_head.borrow_mut().next = Some(old_head);
                self.head = Some(new_head);
            }
            None => {
                self.tail = Some(new_head.clone());
                self.head = Some(new_head);
            }
        }
    }

    pub fn push_back(&mut self, val: T) {
        let new_tail = Node::new(val);
        match self.tail.take() {
            Some(old_tail) => {
                old_tail.borrow_mut().next = Some(new_tail.clone());
                new_tail.borrow_mut().prev = Some(old_tail);
                self.tail = Some(new_tail);
            }
            None => {
                self.head = Some(new_tail.clone());
                self.tail = Some(new_tail);
            }
        }
    }

    pub fn pop_back(&mut self) -> Option<T> {
        self.tail.take().map(|old_tail| {
            match old_tail.borrow_mut().prev.take() {
                Some(new_tail) => {
                    new_tail.borrow_mut().next.take();
                    self.tail = Some(new_tail);
                }
                None => {
                    self.head.take();
                }
            }
            Rc::try_unwrap(old_tail).ok().unwrap().into_inner().val
        })
    }
}

impl<T> Default for LinkedList<T> {
    fn default() -> Self {
        Self::new()
    }
}

#[cfg(test)]
mod tests {
    use std::{
        ops::DerefMut,
        sync::{Arc, Mutex},
        thread,
    };

    use super::*;

    #[test]
    fn it_works() {
        let node = Node::new(1);
        let val = node.borrow().val;
        assert_eq!(val, node.borrow().val);
    }

    #[test]
    fn sync_test() {
        let state = Arc::new(Mutex::new(String::from("abc")));
        let st1 = Arc::clone(&state);
        let jh1 = thread::spawn(move || {
            st1.lock().unwrap().deref_mut().push_str("123");
        });

        let st2 = Arc::clone(&state);
        let jh2 = thread::spawn(move || {
            st2.lock().unwrap().deref_mut().push_str("456");
        });

        jh1.join().unwrap();
        jh2.join().unwrap();

        let s = state.lock();
        let s = s.as_deref().unwrap();
        println!("{}", s);
    }
}
