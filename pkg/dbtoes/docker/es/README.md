```shell
# docker run 
docker run -d --name elasticsearch --network dbtoes -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" --name="es"  elasticsearch:7.17.13
# 拷贝配置到本地目录,目录从workdir找 或者 去hub.docker找
docker cp es:/usr/share/elasticsearch/config ./



# port 
9200,9300
# workdir
/usr/share/elasticsearch
# Entrypoint
 [
    "/bin/tini",
    "--",
    "/usr/local/bin/docker-entrypoint.sh"
]


```