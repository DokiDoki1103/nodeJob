# 基于 nginx 轻量版镜像构建
FROM nginx:alpine

# 复制自定义 nginx.conf 文件
COPY nginx.conf /etc/nginx/conf.d/default.conf

# 暴露端口
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
