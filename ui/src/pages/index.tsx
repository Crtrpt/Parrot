import React, { useEffect } from 'react';
import { LaptopOutlined, NotificationOutlined, UserOutlined } from '@ant-design/icons';
import type { MenuProps } from 'antd';
import { Breadcrumb, Layout, Menu, theme } from 'antd';

import { useState } from 'react';


import {GetMenuList} from "../api"

const { Header, Content, Sider } = Layout;

let items1: MenuProps['items'] = ['1', '2', '3'].map((key) => ({
  key,
  label: `nav ${key}`,
}));

const rightNav: MenuProps['items'] = [
  {
    key: `1`,
    label: `个人信息`,
  },
  {
    key: `2`,
    label: `设置`,
  },
  {
    key: `3`,
    label: `退出`,
  },
  
]


const items2: MenuProps['items'] = [UserOutlined, LaptopOutlined, NotificationOutlined].map(
  (icon, index) => {
    const key = String(index + 1);

    return {
      key: `sub${key}`,
      icon: React.createElement(icon),
      label: `subnav ${key}`,
    };
  },
);

const Index: React.FC = () => {
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();


  const [nav, setNav] = useState(items1);
 
  useEffect(()=>{
    GetMenuList().then((resp)=>{
      setNav(resp.Data)

      console.log(resp.Data)
    })
  },[])

 

  return (
    <Layout>
      <Header style={{ display: 'flex' ,alignItems:"flex-start"}}>
        <div  style={{ display: 'flex',flexGrow:"1" }}>
        <div className="demo-logo" />
          <Menu
            theme="dark"
            mode="horizontal"
            defaultSelectedKeys={['2']}
            items={nav}
            style={{ flex: 1, minWidth: 0 }}
          />
        </div>
   
      <div  style={{ display: 'flex' }}>
      <Menu
          theme="dark"
          mode="horizontal"
          defaultSelectedKeys={['2']}
          items={rightNav}
          style={{ flex: 1, minWidth: 0 }}
        />
      </div>

      </Header>
      <Layout>
        <Sider width={200} style={{ background: colorBgContainer }}>
          <Menu
            mode="inline"
            defaultSelectedKeys={['1']}
            defaultOpenKeys={['sub1']}
            style={{ height: '100%', borderRight: 0 }}
            items={items2}
          />
        </Sider>
        <Layout style={{ padding: '0 24px 24px' }}>
          <Breadcrumb
            items={[{ title: 'Home' }, { title: 'List' }, { title: 'App' }]}
            style={{ margin: '16px 0' }}
          />
          <Content
            style={{
              padding: 24,
              margin: 0,
              minHeight: 280,
              background: colorBgContainer,
              borderRadius: borderRadiusLG,
            }}
          >
            Content
          </Content>
        </Layout>
      </Layout>
    </Layout>
  );
};

export default Index;