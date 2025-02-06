#!/bin/bash

# 函数用于打印带有颜色的日志信息
log_info() {
    echo -e "\033[32m[INFO] $1\033[0m"
}

log_warn() {
    echo -e "\033[33m[WARN] $1\033[0m"
}

log_error() {
    echo -e "\033[31m[ERROR] $1\033[0m"
}

# 检查是否提供了分支名作为参数
if [ $# -eq 0 ]; then
    log_error "请提供要推送的分支名作为参数，例如：$0 1.1"
    exit 1
fi

BRANCH=$1
# 最大重试次数
# MAX_ATTEMPTS=10
# 当前尝试次数，初始为 0
ATTEMPT=0

log_info "开始尝试将分支 $BRANCH 推送到远程仓库..."

# while [ $ATTEMPT -lt $MAX_ATTEMPTS ]; do
while true; do
    log_info "第 $((ATTEMPT + 1)) 次尝试推送分支 $BRANCH 到远程仓库..."
    # 执行 git push 命令
    git push origin "$BRANCH"
    # 获取上一个命令的退出状态码
    STATUS=$?
    
    if [ $STATUS -eq 0 ]; then
        # 如果状态码为 0，表示命令执行成功
        log_info "Git push 成功，共尝试 $((ATTEMPT + 1)) 次。"
        break
    else
        # 命令执行失败，增加尝试次数
        ATTEMPT=$((ATTEMPT + 1))
        log_warn "Git push 失败。第 $ATTEMPT 次尝试，共 $MAX_ATTEMPTS 次尝试。将在 5 秒后重试..."
        # 等待 5 秒后重试
        sleep 5
    fi
done

# if [ $ATTEMPT -eq $MAX_ATTEMPTS ]; then
#     # 如果达到最大尝试次数仍未成功，输出失败信息
#     log_error "经过 $MAX_ATTEMPTS 次尝试后，Git push 仍然失败。请检查你的网络或仓库。"
# fi