use sudoku_solver::puzzle::Puzzle;

fn main() {
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
    let mut puzz = Puzzle::new(input).unwrap();
    puzz.solve(0);
    let s = puzz.is_solved();
    let res: Vec<Vec<isize>> = puzz.into();
}
