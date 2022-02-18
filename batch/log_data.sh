# 按行读取文件 data.log
# 每行按空格分组，打印第7个元素
# 按字母排序
# 去除重复的行 并统计次数
# 按次数排序，打印前10行

cat data.log | awk '{print $7}' | sort | uniq -c | sort -nr | head -n 10