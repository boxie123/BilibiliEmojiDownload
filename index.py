import os
import httpx
import json
from typing import Dict, List

def getEmojiInfo(client: httpx.Client, id: int) -> Dict[str, str]:
    '''
    获取单个id对应的表情包名称
    '''
    resp = client.get(f"https://api.bilibili.com/bapis/main.community.interface.emote.EmoteService/PackageDetail?id={id}")
    resp.encoding = "UTF-8"
    try:
        emojiInfo = resp.json()
    except json.decoder.JSONDecodeError:
        print(f"id: {id} 请求出错，无法解析")
        return
    if emojiInfo["code"] != 0:
        print(f"id: {id} 请求出错：{emojiInfo}")
        return
    if len(emojiInfo["data"]) < 1:
        print(f"id: {id} 无对应表情包")
        return
    emojiName: str = emojiInfo["data"]["package"]["text"]
    meta: Dict = emojiInfo["data"]["package"]["meta"]
    item: str = meta.get("item_url") or meta.get("item_id")
    return {"id": id, "name": emojiName, "item": item}


if __name__ == "__main__":
    client = httpx.Client(headers={"User-Agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 18_1_1 like Mac OS X) AppleWebKit/619.2.8.10.9 (KHTML, like Gecko) Mobile/22B91 BiliApp/83000100 os/ios model/iPhone 13 mobi_app/iphone build/83000100 osVer/18.1.1 network/2 channel/AppStore Buvid/YF4BDFF823E8BA68449892FA07B6F4028355 c_locale/zh-Hans_CN s_locale/zh-Hans_CN sessionID/11bb9479 disable_rcmd/0"})
    if os.path.exists("index.json"):
        with open("index.json", "r", encoding="utf-8") as f:
            AllInfo: List = json.load(f)
        assert isinstance(AllInfo, list)
    else:
        AllInfo = []
    if len(AllInfo) >= 1:
        start = AllInfo[-1]["id"]
    else:
        start = 0
    end = max(start + 100, 7650)

    for i in range(start + 1, end + 1):
        emojiInfo = getEmojiInfo(client, i)
        if emojiInfo:
            AllInfo.append(emojiInfo)
            print(f"id: {i} 获取成功")
    with open("index.json", "w", encoding="utf-8") as f:
        json.dump(AllInfo, f, ensure_ascii=False, indent=2)