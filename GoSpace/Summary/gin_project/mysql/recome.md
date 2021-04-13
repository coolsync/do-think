COALESCE:聚合

在本教程中，您将了解SQL聚合函数，包括：`AVG()`，`COUNT()`，`MIN()`，`MAX()`和`SUM()`。

SQL聚合函数计算一组值并返回单个值。 例如，平均函数(`AVG`)采用值列表并返回平均值。

因为聚合函数对一组值进行操作，所以它通常与`SELECT`语句的[GROUP BY](http://www.yiibai.com/sql/sql-group-by.html)子句一起使用。 `GROUP BY`子句将结果集划分为值分组，聚合函数为每个分组返回单个值。

```sql
SELECT c1, aggregate_function(c2)
FROM table
GROUP BY c1;
SQL
```

以下是常用的SQL聚合函数：

- [AVG()](http://www.yiibai.com/sql/sql-avg.html) - 返回集合的平均值。
- [COUNT()](http://www.yiibai.com/sql/sql-count.html) - 返回集合中的项目数。
- [MAX()](http://www.yiibai.com/sql/sql-max.html) - 返回集合中的最大值。
- [MIN()](http://www.yiibai.com/sql/sql-min.html) - 返回集合中的最小值
- [SUM()](http://www.yiibai.com/sql/sql-sum.html) - 返回集合中所有或不同值的总和。

除`COUNT()`函数外，SQL聚合函数忽略`null`值。只能将聚合函数用作表达式，如下所示：

- `SELECT`语句的选择列表，子查询或外部查询。
- 一个[HAVING子句](http://www.yiibai.com/sql/sql-having.html)



