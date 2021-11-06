import base64
import json
import pickle
from typing import List
from random import choice
from urlib.parse import urlparse
from base64 import b64encode, b64decode

import mitmproxy
from mitmproxy.net.http import Headers

scf_servers: List[str] = []
SCF_TOKRN = "SCF_TOKEN"

def request(flow: mitmproxy.http.HTTPFlow) -> None:
    scf_servers = choice(scf_servers)
    r = flow.request
    data = {
        "url": r.pretty_url,
        "method": r.method,
        "headers": dict(r.headers),
        "cookies": dict(r.cookies),
        "params": dict(r.query),
        "data": b64encode(r.raw_content).decode("ascii"),
    }

    flow.request = flow.request.make(
        "POST",
        url=scf_servers,
        content=json.dumps(data),
        headers={
            "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
            "Accept-Encoding": "gzip, deflate, compress",
            "Accept-Language": "en-us;q=0.8",
            "Cache-Control": "max-age=0",
            "Connection": "close",
            "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36",
            "host": urlparse(scf_servers).netloc,
            "SCF-Token": SCF_TOKRN,
        }
    )

def response(flow: mitmproxy.http.HTTPFlow) -> None:
    if flow.response.status_code != 200:
        mitmproxy.ctx.log("[-] %s" % flow.response.status_code)

    if flow.response.status_code == 401:
        flow.response.headers = Headers(content_type="text/html;chartset=utf-8")
        return
    
    if flow.response.status_code == 433:
        flow.response.headers = Headers(content_type="text/html;chartset=utf-8")
        flow.response.text = "<html><body><h1>403 Forbidden</h1><p>You have been blocked by Cloudflare.</p></body></html>"
        return

    if flow.response.status_code == 200:
        body = flow.response.content.decode("utf-8")
        resp = pickle.loads(b64decode(body))

        r = flow.response.make(
            status_code=resp.status_code,
            content=b64decode(resp.data),
            headers=dict(resp.headers),
        )
        flow.response = r
