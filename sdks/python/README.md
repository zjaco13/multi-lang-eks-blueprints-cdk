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
git clone https://github.com/zjaco13/multi-lang-eks-blueprints-cdk.git
cd multi-lang-eks-blueprints-cdk
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

Add this to the `main.py` file
```python
from eks_blueprints_python_sdk import builder

def build(stub):
    ...

if __name__ == "__main__":
   builder.run(build) 

```

Add function calls to the gRPC server in the build function to build your EKS Blueprint

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
make deploy-server
```

## Contributing

Publishing

A new package will need to be released whenever protobufs change.

First make an account on [test.pypi.org](https://test.pypi.org/), and you must be added to the [eks-blueprints-python-sdk project](https://test.pypi.org/project/eks-blueprints-python-sdk/)

This project uses the [setuptools](https://setuptools.pypa.io/en/latest/) build backend.  If not installed, install it with this command

```bash
pip install setuptools
```

You will also need the build package installed

```bash
pip install build
```

Run the build with this command

```bash
python -m build
```

This will create a `dist` folder under `multi-lang-eks-blueprints-cdk/sdks/python`

You will need an api token to upload the package to TestPyPI, and you will need to set up a `.pypirc` in your home directory similar to this

```sh
[distutils]
index-servers =
    testpypi

[testpypi]
username = __token__
password = <<YOUR_API_TOKEN>>
```

To upload the package, install twine

```bash
pip install twine
```

Run twine to upload all of the archives under dist

```bash
python -m twine upload --repository testpypi dist/*
```

Changes to the packaging process can be made by updating the `pyproject.toml` and `MANIFEST.in` files.

