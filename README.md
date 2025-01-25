# bilibiliemojidownload

爬取b站装扮表情包（自用）

## 用法

### 搭建 python 环境并生成索引

> [!NOTE]
> 本项目中包含`index.json`文件（于2025/01/25生成），若其中包含你想要的表情包信息，则可忽视本小节，并跳至[golang部分](#运行-golang-部分)

1. 需要当前用户已正确安装`python`，版本>=3.8
2. `clone`本项目：`git clone https://github.com/boxie123/BilibiliEmojiDownload.git`并进入项目文件夹`cd BilibiliEmojiDownload`
3. 若已安装[`rye`](https://github.com/astral-sh/rye)：

    ```batch
    rye sync
    rye run python .\index.py
    ```

    若未安装：

    ```batch
    python -m venv .venv
    .venv\Scripts\activate
    pip install httpx
    python .\index.py
    ```

### 运行 golang 部分

1. 打开上一步生成的 `index.json`，找到你想要的表情包的 `id`（若没有`id`，则需根据`url`自行下载，本程序无法直接爬取）；
2. 下载`release`中的`exe`文件并运行，或直接运行：`go run ./main.go`；
3. 输入第2步中获取的id，并回车运行，等待下载结束。

## 声明

> [!CAUTION]
> 请勿滥用，本项目仅用于学习和测试！请勿滥用，本项目仅用于学习和测试！请勿滥用，本项目仅用于学习和测试！
>
> 本项目遵守 CC-BY-NC 4.0 协议，禁止一切商业使用，造成的一切不良后果与本人无关！
