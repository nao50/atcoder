fn solve(h1: i32, h2: i32) -> i32 {
    return h1 - h2;
}

fn parse_int() -> i32 {
    let mut l = String::new();
    std::io::stdin()
        .read_line(&mut l)
        .expect("Failed to read line");
    return l.trim().parse().expect("");
}

fn main() {
    let h1 = parse_int();
    let h2 = parse_int();
    print!("{}", solve(h1, h2));
}

#[test]
fn solve_test() {
    assert_eq!(solve(0, 0), 0);
    assert_eq!(solve(0, 1), -1);
    assert_eq!(solve(1, 0), 1);
}
