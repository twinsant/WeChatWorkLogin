# WeChatWorkLogin
 企业微信登录

 ## React
 ```Typescript
 import { Row, Col } from 'antd';
import React, { Component } from 'react';
import Script from 'react-load-script';
import consola from 'consola';

class Login extends Component {
  handleScriptLoad() {
    consola.info(MY_ENV);
    window.WwLogin({
      id: 'wx_reg',
      appid: WECHAT_WORK_APPID,
      agentid: WECHAT_WORK_AGENTID,
      redirect_uri: WECHAT_WORK_REDIRECT_URI,
      state: '',
      href: '',
    });
  }
  render() {
    return (
      <div>
        <Row type="flex" justify="center" align="top">
          <div id="wx_reg" />
          <Script
            url="http://rescdn.qqmail.com/node/ww/wwopenmng/js/sso/wwLogin-1.0.0.js"
            onLoad={this.handleScriptLoad.bind(this)}
          />
        </Row>
      </div>
    );
  }
}

export default Login;

 ```

 ## Vue + Element
 ```Javascript
 <template>
  <el-row type="flex" justify="center" class="login-container">
    <el-col :span="6">
      <div id="wx_reg" class="grid-content">
      </div>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: 'login',
  data() {
    return {
    }
  },
  mounted() {
    const _this = this;
    this.loadScript('http://rescdn.qqmail.com/node/ww/wwopenmng/js/sso/wwLogin-1.0.0.js', r => {
      window.WwLogin({
              "id" : "wx_reg",  
              "appid" : process.env.WECHAT_WORK_APPID,
              "agentid" : process.env.WECHAT_WORK_AGENTID,
              "redirect_uri" : process.env.WECHAT_WORK_REDIRECT_URI,
              "state" : "",
              "href" : "",
      });
    })
  },
  methods: {
    loadScript(src, callback) {
      let script = document.createElement('script');
      script.type = 'text/javascript';
      if (callback) script.onload = callback;
      document.querySelector('body').appendChild(script);
      script.src = src;
    },
  },
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
$bg:#2d3a4b;

.login-container {
  position: fixed;
  height: 100%;
  width: 100%;
  background-color: $bg;
  padding-top: 70px;
}
</style>
```

## Nginx
```
    server {
        listen 80;
        server_name yourdomain;

        location / {
            proxy_pass http://127.0.0.1:9527;
        }
        location /user {
            proxy_pass http://127.0.0.1:8080;
        }
    }
```