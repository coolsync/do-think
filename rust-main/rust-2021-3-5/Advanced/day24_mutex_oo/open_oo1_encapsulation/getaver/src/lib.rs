pub struct AverCollection {
    list: Vec<i32>,
    aver: f64,
}

impl AverCollection {
    pub fn new() -> AverCollection {
        AverCollection {
            list: vec![],
            aver: 0.0,
        }
    }

    pub fn add(&mut self, value: i32) {
        self.list.push(value);
        self.update_aver();
    }

    pub fn remove(&mut self) -> Option<i32> {
        let result = self.list.pop();
        match result {
            Some(value) => {
                self.update_aver();
                Some(value)
            },
            None => None,
        }
    }

    pub fn average(&self) -> f64 {
        self.aver
    }
    
    fn update_aver(&mut self) {
        let total: i32 = self.list.iter().sum();
        self.aver = total as f64 / self.list.len() as f64;
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}
