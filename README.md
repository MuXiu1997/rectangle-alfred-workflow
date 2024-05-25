# Rectangle-Alfred-Workflow

使用 [`Alfred`](https://www.alfredapp.com/) + 拼音来调用 [`Rectangle`](https://github.com/rxhanson/Rectangle), 避免设置过多快捷键




https://user-images.githubusercontent.com/49554020/182882290-efcf33cc-4aa1-4c23-81f9-b368a26d5c0e.mp4



## 实现原理

~~1. 将 `Rectangle` 的所有快捷键清空~~

~~2. 实现一个 `filter` 将 `Alfred` 中输入的拼音转化为 `Rectangle` 的动作 ID~~

~~3. 使用 `defaults write` 为该动作设置一个不常用的快捷键~~

~~4. 使用 `Alfred` 调用该快捷键~~

~~5. 再次使用 `defaults write` 清空该动作的快捷键~~

参考: [Rectangle - Execute an action by URL](https://github.com/rxhanson/Rectangle#execute-an-action-by-url)



## Debug

env:

```
alfred_workflow_bundleid=net.deanishe.awgo;alfred_workflow_data=./.temp/testenv/data;alfred_workflow_cache=./.temp/testenv/cache;alfred_version=3.8.1;alfred_workflow_version=1.2.0;alfred_workflow_name=AwGo;AW_SESSION_ID=test-session-id
```
