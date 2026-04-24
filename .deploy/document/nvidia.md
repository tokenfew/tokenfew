# MOTHERBOARD

⚡ Configuration instructions for the motherboard. ⚡

## License

[![License:Apache2.0](https://img.shields.io/badge/License-Apache2.0-yellow.svg)](https://opensource.org/licenses/Apache2.0)

## Basic Requirements

Your device should meet the following basic requirements.

```shell
Distributor ID: Ubuntu
Description:    Ubuntu 24.04.4 LTS
Release:        24.04
Codename:       noble
```

## General command script

Below is a generic installation script

```bash
sudo apt install -y build-essential linux-headers-$(uname -r) dkms htop nvtop neofetch
```

```bash
# Check whether the graphics card is mounted
lspci I grep -i nvidia
```

### Tesla V100-SXM2-16GB/32G

Suitable for Tesla V100-SXM2-16GB/32G server graphics card

#### Step 1：Obtain the driver installation package

Go to `nvidia.cn/drivers` to obtain the `deb` installation package for the corresponding version below

```shell
Driver Version: 570.211.01
CUDA Version: 12.8
```

Usually, you will get an installation file like `nvidia-driver-local-repo-ubuntu2404-570.211.01_1.0-1_amd64.deb`

#### Step 2：Install deb files

Use the following command to install the `deb` file

```bash
sudo dpkg -i nvidia-driver-local-repo-ubuntu2404-570.211.01_1.0-1_amd64.deb
```

```bash
sudo cp /var/nvidia-driver-local-repo-ubuntu2404-570.211.01/nvidia-driver-local-0E456948-keyring.gpg /usr/share/keyrings/
```

```bash
sudo apt update
```

#### Step 2：Install driver and cuda toolkit

Use the following command to install the driver

```bash
sudo apt install -y nvidia-driver-570
```

(Optional) Use the following command to install the CUDA toolkit

```bash
wget https://developer.download.nvidia.com/compute/cuda/repos/ubuntu2404/x86_64/cuda-keyring_1.1-1_all.deb
sudo dpkg -i cuda-keyring_1.1-1_all.deb
sudo apt update
```

```bash
sudo apt install -y cuda-toolkit-12-8
```

```bash
# version verification
/usr/local/cuda-12.8/bin/nvcc -V
```

#### Step 3：verification

You can verify whether the driver has been installed by executing the following command

```bash
nvidia-smi
```

```bash
nvidia-smi nvlink -s
```
