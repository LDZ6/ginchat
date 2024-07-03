## Git 使用详解

### 添加和提交文件

#### 添加文件到暂存区
```bash
git add .            # 添加所有文件到暂存区
git add <文件名>     # 添加指定文件到暂存区
```

#### 提交到本地仓库
```bash
git commit -m "提交说明"  # 提交到本地仓库，并附上提交说明
```

### 查看状态和历史

#### 查看状态
```bash
git status           # 查看当前工作区状态
```

#### 查看修改内容
```bash
git diff             # 查看未暂存的修改内容
```

#### 查看提交历史
```bash
git log              # 查看提交历史记录
```

#### 查看提交历史详细信息
```bash
git relog -p         # 查看提交历史的详细信息
```

### 版本回退

#### 回退到上一个版本
```bash
git reset --hard HEAD^   # 回退到上一个版本
```

#### 回退到指定版本
```bash
git reset --hard <版本号> # 回退到指定版本
```

### 配置 .gitignore 文件

#### 创建 .gitignore 文件
```bash
touch .gitignore         # 创建 .gitignore 文件
```

#### .gitignore 文件配置示例
```plaintext
*.css           # 忽略所有 .css 文件
!important.css  # 不忽略 important.css 文件
asdf.css        # 忽略 asdf.css 文件
```

### 推送到远程仓库

#### 推送本地分支到远程仓库
```bash
git push origin master   # 将本地 master 分支推送到远程仓库
```

### 克隆远程仓库

#### 克隆远程仓库到本地
```bash
git clone <仓库地址>     # 克隆远程仓库到本地
```

### 配置用户信息

#### 设置用户名
```bash
git config user.name "用户名"   # 设置 Git 用户名
```

#### 设置邮箱
```bash
git config user.email "邮箱"   # 设置 Git 邮箱
```

#### 查看配置信息
```bash
git config -l              # 查看 Git 的配置信息
```

### Go 模块初始化

#### 初始化 Go 模块
选择模块路径，通常基于您的域名。例如，如果您的 GitHub 用户名是 `LDZ6` 并且您的项目名是 `ginchat`，可以使用 `github.com/LDZ6/ginchat` 作为模块路径。

```bash
go mod init github.com/LDZ6/ginchat  # 初始化 Go 模块
```
下面是详细的Markdown笔记，涵盖了多人开发过程中使用Git的基本步骤和解决冲突的方法。

### 多人开发使用Git详解

#### 1. 设置服务器端仓库

在服务器上初始化一个裸仓库（Bare Repository）用于共享代码：

```bash
git init --bare <仓库名称>.git
```

#### 2. 开发人员克隆远程仓库

开发人员可以克隆服务器上的仓库到本地进行开发：

```bash
git clone <远程仓库地址>
cd <仓库名称>
```

#### 3. 开发阶段的常用操作

在开发阶段，开发人员需要配置个人信息，并进行代码的提交和同步操作：

```bash
git config user.name "用户名"
git config user.email "邮箱"

# 添加所有修改的文件
git add .

# 提交更改
git commit -m "提交说明"

# 查看仓库当前状态
git status

# 推送本地提交到远程仓库
git push

# 如果是第一次推送，设置本地分支与远程分支的关联
git push --set-upstream origin master
```

#### 4. 多人开发与单人开发的区别

在多人开发中，除了每个人都可以向远程仓库推送更改外，还需要及时拉取其他开发者的更改：

```bash
# 拉取远程仓库的更新
git pull
```

#### 5. 处理代码冲突

如果多个开发者同时修改了同一文件，并且提交到远程仓库后产生冲突，需要解决冲突后再提交：

- 当开发者B在A之后提交时可能遇到冲突，解决冲突的步骤如下：

```bash
# 拉取远程仓库的更新
git pull

# 解决冲突后，手动编辑文件以解决冲突
# 再次添加、提交更改
git add .
git commit -m "解决冲突"

# 推送解决后的更改
git push
```

之后A进行pull,是B解决冲突后的提交，因为B在处理冲突时已经合并了A的更改。

这些步骤可以帮助团队有效地进行协作开发，确保代码的同步和质量。