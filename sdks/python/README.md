# Python Client for Multi Language EKS Blueprints for cdk

## Run the example
To run the example, first start the server from https://github.com/aws-quickstart/cdk-eks-blueprints/tree/blueprints-api-support
```bash
git clone https://github.com/aws-quickstart/cdk-eks-blueprints.git
cd cdk-eks-blueprints
git checkout blueprints-api-support
make run-server
```

Then run the python example by cloning this repository and running with make, a virtual env with grpcio and protobuf will be created
```bash
make python-example
```

## How to use as a package

Create a new python file 
```bash
touch main.py
```

Optionally create a new virtual environment and activate it
```bash
python3 -m venv venv
source venv/bin/activate
```

Install the sdk and other dependencies
```bash
pip install --index-url https://test.pypi.org/simple/ --no-deps eks-blueprints-python-sdk
pip install grpcio
pip install protobuf
```

Add this to the main.py file
```python
from eks_blueprints_python_sdk import builder

def build(stub):
    ...

if __name__ == "__main__":
   builder.run(build) 

```

Add function calls to the grpc server in the build function to build your EKS Blueprint

Run the server
```bash
git clone https://github.com/aws-quickstart/cdk-eks-blueprints.git
cd cdk-eks-blueprints
git checkout blueprints-api-support
make run-server
```

Run your code
```bash
python3 main.py
```

From the server, deploy your EKS Blueprint
```bash
npx cdk -a cdk.out/ deploy
```

## Contributing

To compile protobufs to python you will need grpcio-tools installed
```bash
pip install grpcio-tools
```



