# DETL : Distributed ETL

* https://github.com/swafran/detl-common
* https://github.com/swafran/detl-transform

Usage recommendation: do any aggregation or data joins during extraction. Most of the time extraction calls an API with SQL behind it, SQL directly, or static documents. The tools used naturally lend themselves to aggregation. (If you really do want aggregation during transform, it is possible with a custom handler).