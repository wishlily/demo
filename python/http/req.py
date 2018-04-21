import requests

def getHTMLText(url):
	try:
		# kv = {'user-agent':'Mozilla/5.0'} headers=kv
		# kv = {'wd':'Python'} params = kv ?wd=Python
		r = requests.get(url, timeout=30)
		r.raise_for_status() # 如果状态不是 200，引发异常
		r.encoding = r.apparent_encoding
		return r.text
	except:
		return "ERROR"

if __name__ == "__main__":
	url = "https://www.baidu.com"
	print(getHTMLText(url))