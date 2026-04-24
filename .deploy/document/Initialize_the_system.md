# Ubuntu 24\.04\.04 完整一键安装 NVIDIA 570\.211\.01 驱动教程（含deb包\+run包双方案）

下面给你 **Ubuntu 24\.04\.04 完整一键流程**，包含两种安装方式：官方闭源 `\.run` 驱动 \+ 你当前使用的 `\.deb` 驱动，均包含**换阿里云源、清旧驱动、禁 nouveau、关桌面、内核依赖、验证、恢复桌面**全步骤，直接复制命令执行即可。

> 570\.211\.01 适配 Ubuntu 24\.04 默认 **6\.8 内核**，AI 大模型、CUDA、vLLM 部署通用；**必须先关 Secure Boot**，否则内核模块加载失败；换阿里云源可解决后续依赖安装缓慢、失败问题。

## 一、前置准备：替换 Ubuntu 24\.04 阿里云官方源

替换为阿里云源，提升系统更新、依赖安装速度，避免后续驱动安装因源问题失败，步骤如下：

### 1\. 备份原始源（防止替换出错，可回滚）

```bash
sudo cp /etc/apt/sources.list /etc/apt/sources.list.bak.$(date +%Y%m%d)
```

### 2\. 写入完整阿里云 noble\(24\.04\) 专用源

```bash
sudo tee /etc/apt/sources.list <<EOF
deb http://mirrors.aliyun.com/ubuntu/ noble main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ noble-updates main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ noble-backports main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ noble-security main restricted universe multiverse
EOF
```

### 3\. 更新源缓存（使新源生效）

```bash
sudo apt update
```

## 二、安装前必做准备（两种安装方式通用）

### 1\. 关闭 BIOS Secure Boot

重启电脑进 BIOS/UEFI，找到 **Secure Boot** 设置为 **Disabled**，保存退出。

> Ubuntu 24\.04 默认开启安全启动，会拦截第三方 NVIDIA 内核模块，`nvidia\-smi` 永远报错。

### 2\. 更新系统 \&amp; 安装编译依赖（内核头文件、gcc、make）

```bash
sudo apt update && sudo apt upgrade -y
sudo apt install -y build-essential linux-headers-$(uname -r) dkms gcc make
```

### 3\. 彻底卸载所有旧 NVIDIA 驱动残留

```bash
# 清理系统apt安装的所有nvidia包
sudo apt purge -y *nvidia* *cuda*
sudo apt autoremove -y && sudo apt autoclean
# 若之前装过run驱动，执行卸载
sudo nvidia-uninstall
```

### 4\. 永久禁用系统自带 nouveau 开源驱动（核心步骤）

一键写入黑名单配置并更新内核镜像：

```bash
sudo tee /etc/modprobe.d/blacklist-nouveau.conf << EOF
blacklist nouveau
blacklist lbm-nouveau
options nouveau modeset=0
alias nouveau off
alias lbm-nouveau off
EOF

# 更新initramfs，让配置重启生效
sudo update-initramfs -u
```

### 5\. 重启系统

```bash
sudo reboot
```

### 6\. 重启后验证 nouveau 已禁用

执行下面命令**无任何输出**即为成功：

```bash
lsmod | grep nouveau
```

## 三、方案一：\.deb 包安装（你当前使用的方式，解决GPG密钥问题）

针对你下载的 `nvidia\-driver\-local\-repo\-ubuntu2404\-570\.211\.01\_1\.0\-1\_amd64\.deb`，解决GPG密钥缺失问题，完整步骤如下：

### 1\. 进入deb包所在目录（你的包在“下载”目录）

```bash
cd ~/下载
```

### 2\. 赋予文件可执行权限

```bash
chmod +x nvidia-driver-local-repo-ubuntu2404-570.211.01_1.0-1_amd64.deb
```

### 3\. 安装deb包

```bash
sudo dpkg -i nvidia-driver-local-repo-ubuntu2404-570.211.01_1.0-1_amd64.deb
```

### 4\. 解决GPG密钥缺失问题（关键步骤）

根据终端提示，执行以下命令安装缺失的GPG密钥，否则后续无法更新和安装驱动依赖：

```bash
sudo cp /var/nvidia-driver-local-repo-ubuntu2404-570.211.01/nvidia-driver-local-0E456948-keyring.gpg /usr/share/keyrings/
```

### 5\. 更新源缓存，使deb包仓库生效

```bash
sudo apt update
```

### 6\. 安装NVIDIA驱动（自动安装依赖）

```bash
sudo apt install -y nvidia-driver-570
```

安装过程中无需手动干预，等待终端提示安装完成即可。

## 四、方案二：\.run 包安装（原教程方式，备用）

### 1\. 进入纯文本 TTY（关闭图形桌面，防止显卡占用）

Ubuntu 24\.04 GNOME 桌面会占用 GPU，**必须关闭桌面服务才能装 run 驱动**：

```bash
# 停止gdm3图形管理器，进入命令行模式
sudo systemctl stop gdm3
# 切换到多用户文本模式
sudo systemctl isolate multi-user.target
```

此时桌面黑屏，仅保留终端登录界面。

### 2\. 进入驱动文件所在目录

把下载好的 `NVIDIA\-Linux\-x86\_64\-570\.211\.01\.run` 放到用户目录，示例：

```bash
cd ~
```

### 3\. 赋予文件可执行权限

```bash
chmod +x NVIDIA-Linux-x86_64-570.211.01.run
```

### 4\. 执行安装（推荐静默参数，避坑）

```bash
sudo ./NVIDIA-Linux-x86_64-570.211.01.run \
--no-x-check \
--no-nouveau-check \
--dkms \
--silent
```

参数说明：

- `\-\-no\-x\-check`：跳过 X 桌面检测（已关桌面必备）

- `\-\-no\-nouveau\-check`：跳过 nouveau 检测

- `\-\-dkms`：自动生成内核模块，内核更新后驱动自动适配

- `\-\-silent`：静默安装，全程自动无需手动点选

安装完成终端提示 **Installation complete** 即成功。

## 五、验证驱动安装成功（两种方案通用）

```bash
nvidia-smi
```

正常输出显卡信息、驱动版本 **570\.211\.01**、CUDA 版本，即为安装完成。

## 六、恢复图形桌面（\.run包安装必做，\.deb包安装可跳过）

```bash
# 恢复桌面显示管理器
sudo systemctl start gdm3
# 切回图形用户模式
sudo systemctl isolate graphical.target
```

等待几秒，桌面自动恢复。

## 七、常见报错 \&amp; 解决方案

### 报错 1：The Nouveau kernel driver is currently in use

原因：nouveau 未彻底禁用
解决：重新执行**步骤 4 禁用 nouveau** \+ `sudo update\-initramfs \-u` \+ **重启**。

### 报错 2：Unable to load kernel module / Secure Boot enabled

原因：BIOS Secure Boot 开启
解决：重启进 BIOS 关闭**Secure Boot**，再重装驱动。

### 报错 3：gcc version mismatch /kernel headers missing

解决：重新安装内核依赖

```bash
sudo apt install -y linux-headers-$(uname -r) gcc build-essential
```

### 报错 4：nvidia\-smi 命令不存在

原因：安装未完成 / 中途中断
解决：\.deb包安装：重新执行 `sudo apt install \-y nvidia\-driver\-570`；\.run包安装：回到 TTY，重新执行 run 安装命令。

### 报错 5：GPG key does not appear to be installed

原因：deb包安装时，NVIDIA本地仓库的GPG密钥未安装
解决：执行命令 `sudo cp /var/nvidia\-driver\-local\-repo\-ubuntu2404\-570\.211\.01/nvidia\-driver\-local\-0E456948\-keyring\.gpg /usr/share/keyrings/`，再执行 `sudo apt update`。

## 八、后续卸载驱动命令（两种方案通用）

以后需要卸载该版本驱动：

```bash
# 卸载deb包安装的驱动
sudo apt purge -y nvidia-driver-570 nvidia-driver-local-repo-ubuntu2404-570.211.01
# 卸载run包安装的驱动
sudo nvidia-uninstall
# 清理残留
sudo apt autoremove -y && sudo apt autoclean
```

需要我给你配套的 **CUDA 12\.x 对应版本对照表**，以及 Ubuntu 24\.04 安装 CUDA 无冲突脚本吗？

> （注：文档部分内容可能由 AI 生成）
