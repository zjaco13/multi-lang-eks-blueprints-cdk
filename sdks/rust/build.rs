// use glob::glob;

fn main() {
    tonic_build::configure().build_client(true)
        .out_dir("./src")
        .build_client(true)
        .build_server(false)
        .compile(&["../../proto/cluster.proto"], &["../../proto"])
        .unwrap_or_else(|e| panic!("Proto compile error {}", e));
}
