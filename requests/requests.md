---
title: requests.md
date: 2026-05-05 01:05:00
---



接口测试 curl 脚本记录。


## 新增接口

```bash
curl -X POST --json '{
        "content": "取快递",
        "category": "日常生活",
        "isComplete": "N",
        "deadline": "2026-05-05T11:00:00+08:00"
}' http://localhost:8000/todo
```

