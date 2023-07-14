fn main() {
    tonic_build::configure().build_client(true).out_dir("./src/proto").compile(&["../../proto/team.proto"], includes)
}
