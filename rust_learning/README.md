# Rust Til Dust

I learn rust now.
* Most of this repo is <https://github.com/rust-lang/rustlings/>

Create a project: `cargo new projectName`

Comple and run: `cargo run`

Compile a project: `cargo build`
* Alt cmd: `rustc file.rs` 

```none
└──  projectDir
   ├── Cargo.lock #keeps track of deps
   ├── Cargo.toml
   ├── src/
   └── target/* #contains exe
```

Will compile without compiling: `cargo check`

Builtin linter: `cd src; rustfmt main.rs`

## Ownership Rules

1. Each value in Rust has an owner.
1. There can only be one owner at a time.
1. When the owner goes out of scope, the value will be dropped.

## Answers that took me a sec to understand

| File | Code | Note |
| ---  | ---- | ---- |
| variables2.rs | `let x:i32 = 0;` | |
| variables5.rs | `let mut number = 3;` | [Shadowing](https://doc.rust-lang.org/book/ch03-01-variables-and-mutability.html#shadowing) - use `mut` |
| variables6.rs | `const NUMBER:i32 = 3;` | [Constants](https://doc.rust-lang.org/book/ch03-01-variables-and-mutability.html#constants) - always add a type
| functions2.rs | `fn call_me(num:i32){` | Function signature must have type |
| functions4.rs | `fn sale_price(price: u32) -> u32 {` | `-> expected_return_type` ; also u32 since non-negatives only. |
| functions5.rs | `num * num \|\| return num * num;`| Expressions (ret value based on operands) vs statements (ret value is type void)|
| if1.rs        | `if a > b {a} else {b}` | `pub fn` declaration makes any module, function, or data trcuture accessible from inside of external modules.|
| primitive_types3.rs | `let a = [3;300];` | Creates [an array](https://doc.rust-lang.org/stable/book/ch03-02-data-types.html#the-array-type) starting at 3 with size of 300 |
| primitive_types5.rs | ` let cat = ("Furry McFurson", 3.5); let (name,age) = cat;` | Deconstruct a tuple |
| primitive_types6.rs | `let numbers = (1, 2, 3); let second = numbers.1;` | Accessing a tuple |
| move_semantics6.rs | `mut` vs `&` | `mut` when mutating data, ie.e add an element to end of vec. Borrow `&` when you just need read access. | 
| structs2.rs | `let order_template = create_order_template(); let your_order = Order{name: String::from("Hacker in Rust"),count: 1,..order_template};`| `..order_template` to fill in rest that isn't specified. |
| enums2.rs | `enum Message {Move{x:i8,y:i8},ChangeColor(i16,i16,i16),}` | how to enum | 