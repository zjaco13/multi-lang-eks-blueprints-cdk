use glob::glob;

fn main() {
    let paths = glob("../../proto/*.proto").expect("Failed to read glob pattern");
    let path_str: Vec<&str> = Vec::new();
    for path in paths {
        match path {
            Ok(path) => path_str.push(path.to_string_lossy().trim_end()),
            Err(e) => println!("{:?}", e),
        } 
    }
    tonic_build::configure().build_client(true)
        .out_dir("./src/proto")
        .compile(&paths, &["../../proto"])
        .unwrap_or_else(|e| panic!("Proto compile error {}", e));
}
