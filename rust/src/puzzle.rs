use std::collections::HashSet;

#[derive(Debug, PartialEq)]
pub struct Puzzle {
    nums: [Option<u8>; 81],
    rows: Vec<HashSet<u8>>,
    cols: Vec<HashSet<u8>>,
    boxes: Vec<HashSet<u8>>,
}

impl Puzzle {
    pub fn new(input: Vec<Vec<isize>>) -> Result<Self, &'static str> {
        if input.len() != 9 {
            return Err("expected 9 rows");
        }
        let mut puzz = Puzzle {
            nums: [None; 81],
            rows: Vec::new(),
            cols: Vec::new(),
            boxes: Vec::new(),
        };
        for _ in 0..9 {
            puzz.rows.push(HashSet::new());
            puzz.cols.push(HashSet::new());
            puzz.boxes.push(HashSet::new());
        }
        for (i, row) in input.iter().enumerate() {
            if row.len() != 9 {
                return Err("expected 9 columns");
            }
            for (j, n) in row.iter().enumerate() {
                match n {
                    &val @ 1..=9 => {
                        let val = val as u8;
                        puzz.nums[i * 9 + j] = Some(val);
                        puzz.rows[i].insert(val);
                        puzz.cols[j].insert(val);
                        puzz.boxes[3 * (i / 3) + j / 3].insert(val);
                    }
                    0 => {}
                    _ => Err("expected numbers 0-9")?,
                }
            }
        }
        Ok(puzz)
    }
    pub fn solve(&mut self, start: usize) -> Option<()> {
        if start >= self.nums.len() {
            return Some(());
        }
        for (i, entry) in self.nums.clone()[start..].iter().enumerate() {
            match entry {
                Some(_) => {}
                None => {
                    for n in 1..=9 {
                        if self.is_valid_entry(n, start + i) {
                            self.write_at(n, start + i);
                            match self.solve(start + i + 1) {
                                Some(_) => return Some(()),
                                None => self.clear_at(n, start + i),
                            }
                        }
                    }
                    return None;
                }
            };
        }
        Some(())
    }
    fn is_valid_entry(&self, n: u8, idx: usize) -> bool {
        return !self.rows[row_index(idx)].contains(&n)
            && !self.cols[col_index(idx)].contains(&n)
            && !self.boxes[box_index(idx)].contains(&n);
    }
    fn write_at(&mut self, n: u8, idx: usize) {
        self.nums[idx] = Some(n);
        self.rows[row_index(idx)].insert(n);
        self.cols[col_index(idx)].insert(n);
        self.boxes[box_index(idx)].insert(n);
    }
    fn clear_at(&mut self, n: u8, idx: usize) {
        self.nums[idx] = None;
        self.rows[row_index(idx)].remove(&n);
        self.cols[col_index(idx)].remove(&n);
        self.boxes[box_index(idx)].remove(&n);
    }
    pub fn is_solved(&self) -> bool {
        // use fresh hashmaps so we don't mess with internal state
        let mut rows = Vec::with_capacity(9);
        let mut cols = Vec::with_capacity(9);
        let mut boxes = Vec::with_capacity(9);
        for _ in 0..9 {
            rows.push(HashSet::new());
            cols.push(HashSet::new());
            boxes.push(HashSet::new());
        }
        for (i, n) in self.nums.iter().enumerate() {
            match n {
                None => return false,
                _ => {}
            };
            if !rows[row_index(i)].insert(n)
                || !cols[col_index(i)].insert(n)
                || !boxes[box_index(i)].insert(n)
            {
                return false;
            }
        }
        true
    }
}

impl Into<Vec<Vec<isize>>> for Puzzle {
    fn into(self) -> Vec<Vec<isize>> {
        let mut res = Vec::with_capacity(9);
        for i in 0..9 {
            let mut row = Vec::with_capacity(9);
            for n in self.nums[i * 9..i * 9 + 9].iter() {
                match n {
                    Some(n) => row.push(*n as isize),
                    None => row.push(0),
                }
            }
            res.push(row);
        }
        res
    }
}

fn box_index(idx: usize) -> usize {
    3 * (idx / 27) + (idx % 9) / 3
}

fn row_index(idx: usize) -> usize {
    idx / 9
}

fn col_index(idx: usize) -> usize {
    idx % 9
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_indexing() {
        let data = vec![
            (0, 0, 0, 0),
            (1, 0, 1, 0),
            (2, 0, 2, 0),
            (3, 0, 3, 1),
            (4, 0, 4, 1),
            (5, 0, 5, 1),
            (6, 0, 6, 2),
            (7, 0, 7, 2),
            (8, 0, 8, 2),
            (9, 1, 0, 0),
            (10, 1, 1, 0),
            (11, 1, 2, 0),
            (12, 1, 3, 1),
            (13, 1, 4, 1),
            (14, 1, 5, 1),
            (15, 1, 6, 2),
            (16, 1, 7, 2),
            (17, 1, 8, 2),
            (18, 2, 0, 0),
            (19, 2, 1, 0),
            (20, 2, 2, 0),
            (21, 2, 3, 1),
            (22, 2, 4, 1),
            (23, 2, 5, 1),
            (24, 2, 6, 2),
            (25, 2, 7, 2),
            (26, 2, 8, 2),
            (27, 3, 0, 3),
            (28, 3, 1, 3),
            (29, 3, 2, 3),
            (30, 3, 3, 4),
            (31, 3, 4, 4),
            (32, 3, 5, 4),
            (33, 3, 6, 5),
            (34, 3, 7, 5),
            (35, 3, 8, 5),
            (36, 4, 0, 3),
            (37, 4, 1, 3),
            (38, 4, 2, 3),
            (39, 4, 3, 4),
            (40, 4, 4, 4),
            (41, 4, 5, 4),
            (42, 4, 6, 5),
            (43, 4, 7, 5),
            (44, 4, 8, 5),
            (45, 5, 0, 3),
            (46, 5, 1, 3),
            (47, 5, 2, 3),
            (48, 5, 3, 4),
            (49, 5, 4, 4),
            (50, 5, 5, 4),
            (51, 5, 6, 5),
            (52, 5, 7, 5),
            (53, 5, 8, 5),
            (54, 6, 0, 6),
            (55, 6, 1, 6),
            (56, 6, 2, 6),
            (57, 6, 3, 7),
            (58, 6, 4, 7),
            (59, 6, 5, 7),
            (60, 6, 6, 8),
            (61, 6, 7, 8),
            (62, 6, 8, 8),
            (63, 7, 0, 6),
            (64, 7, 1, 6),
            (65, 7, 2, 6),
            (66, 7, 3, 7),
            (67, 7, 4, 7),
            (68, 7, 5, 7),
            (69, 7, 6, 8),
            (70, 7, 7, 8),
            (71, 7, 8, 8),
            (72, 8, 0, 6),
            (73, 8, 1, 6),
            (74, 8, 2, 6),
            (75, 8, 3, 7),
            (76, 8, 4, 7),
            (77, 8, 5, 7),
            (78, 8, 6, 8),
            (79, 8, 7, 8),
            (80, 8, 8, 8),
        ];
        for (idx, exp_row, exp_col, exp_box) in data {
            assert_eq!(exp_row, row_index(idx));
            assert_eq!(exp_col, col_index(idx));
            assert_eq!(exp_box, box_index(idx));
        }
    }

    #[test]
    fn test_new() {
        let input = vec![
            vec![1, 0, 0, 0, 0, 0, 0, 0, 0],
            vec![0, 2, 0, 0, 0, 0, 0, 0, 0],
            vec![0, 0, 3, 0, 0, 0, 0, 0, 0],
            vec![0, 0, 0, 4, 0, 0, 0, 0, 0],
            vec![0, 0, 0, 0, 5, 0, 0, 0, 0],
            vec![0, 0, 0, 0, 0, 6, 0, 0, 0],
            vec![0, 0, 0, 0, 0, 0, 7, 0, 0],
            vec![0, 0, 0, 0, 0, 0, 0, 8, 0],
            vec![0, 0, 0, 0, 0, 0, 0, 0, 9],
        ];
        let puzz = Puzzle::new(input);
        matches!(puzz, Ok(_));
        let puzz = puzz.unwrap();
        for i in 0..9 {
            assert_eq!(true, puzz.rows[i].contains(&(i as u8 + 1)))
        }
        for i in 0..9 {
            assert_eq!(true, puzz.cols[i].contains(&(i as u8 + 1)))
        }
        assert_eq!(true, puzz.boxes[0].contains(&1));
        assert_eq!(true, puzz.boxes[0].contains(&2));
        assert_eq!(true, puzz.boxes[0].contains(&3));
        assert_eq!(true, puzz.boxes[4].contains(&4));
        assert_eq!(true, puzz.boxes[4].contains(&5));
        assert_eq!(true, puzz.boxes[4].contains(&6));
        assert_eq!(true, puzz.boxes[8].contains(&7));
        assert_eq!(true, puzz.boxes[8].contains(&8));
        assert_eq!(true, puzz.boxes[8].contains(&9));
    }

    #[test]
    fn test_new_errs() {
        let input = vec![
            vec![1, 0, 0, 0, 0, 0, 0, 0, 0],
            vec![0, 1, 0, 0, 0, 0, 0, 0, 0],
        ];
        let puzz = Puzzle::new(input);
        assert_eq!(puzz, Err("expected 9 rows"));
        let input = vec![
            vec![1, 0, 0, 0, 0, 0, 0, 0, 0],
            vec![0, 2, 0, 0, 0, 0, 0, 0, 0],
            vec![0, 0, 3, 0, 0, 0, 0, 0, 0],
            vec![0, 0, 0, 4, 0, 0, 0, 0, 0],
            vec![0, 0, 0, 0, 5, 0, 0, 0, 0],
            vec![0, 0, 0, 0, 0, 6, 0, 0, 0],
            vec![0, 0, 0, 0, 0, 0, 7, 0, 0],
            vec![0, 0, 0, 0, 0, 0, 0, 8, 0],
            vec![0, 0, 0, 0, 0, 0, 0, 0],
        ];
        let puzz = Puzzle::new(input);
        assert_eq!(puzz, Err("expected 9 columns"));
        let input = vec![
            vec![1, 0, 0, 0, 0, 0, 0, 0, 0],
            vec![0, 2, 0, 0, 0, 0, 0, 0, 0],
            vec![0, 0, 3, 0, 0, 0, 0, 0, 0],
            vec![0, 0, 0, 4, 0, 0, 0, 0, 0],
            vec![0, 0, 0, 0, 5, 0, 0, 0, 0],
            vec![0, 0, 0, 0, 0, 6, 0, 0, 0],
            vec![0, 0, 0, 0, 0, 0, 7, 0, 0],
            vec![0, 0, 0, 0, 0, 0, 0, 8, 0],
            vec![0, 0, 0, 0, 0, 0, 0, 0, 10],
        ];
        let puzz = Puzzle::new(input);
        assert_eq!(puzz, Err("expected numbers 0-9"));
    }
    #[test]
    fn test_solve() {
        let input = vec![
            vec![0, 0, 9, 0, 1, 6, 0, 4, 2],
            vec![1, 0, 4, 2, 0, 9, 0, 6, 0],
            vec![0, 2, 0, 0, 0, 8, 7, 0, 0],
            vec![3, 5, 0, 0, 9, 0, 1, 0, 0],
            vec![0, 6, 7, 4, 0, 1, 9, 0, 5],
            vec![0, 0, 0, 7, 5, 0, 0, 8, 6],
            vec![0, 9, 0, 0, 0, 4, 8, 5, 7],
            vec![8, 0, 0, 9, 6, 0, 0, 2, 0],
            vec![4, 7, 0, 8, 0, 5, 0, 0, 0],
        ];
        let solved = vec![
            vec![7, 8, 9, 5, 1, 6, 3, 4, 2],
            vec![1, 3, 4, 2, 7, 9, 5, 6, 8],
            vec![5, 2, 6, 3, 4, 8, 7, 1, 9],
            vec![3, 5, 8, 6, 9, 2, 1, 7, 4],
            vec![2, 6, 7, 4, 8, 1, 9, 3, 5],
            vec![9, 4, 1, 7, 5, 3, 2, 8, 6],
            vec![6, 9, 2, 1, 3, 4, 8, 5, 7],
            vec![8, 1, 5, 9, 6, 7, 4, 2, 3],
            vec![4, 7, 3, 8, 2, 5, 6, 9, 1],
        ];
        let mut puzz = Puzzle::new(input).unwrap();
        matches!(puzz.solve(0), Some(_));
        assert_eq!(true, puzz.is_solved());
        let actual: Vec<Vec<isize>> = puzz.into();
        assert_eq!(solved, actual);
    }
}
