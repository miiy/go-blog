import React from 'react'
import { Menu } from 'antd'
import { Link, NavLink } from 'react-router-dom'

export default class Nav extends React.Component {
  render() {
    return(
      <Menu theme="dark" mode="horizontal" defaultSelectedKeys={['2']}>
        <Menu.Item key="1">
            <NavLink to="/">首页</NavLink>
        </Menu.Item>
        <Menu.Item key="2">
          <NavLink to="/articles">文章</NavLink>
        </Menu.Item>
        <Menu.Item key="3">
          <NavLink to="/about">关于</NavLink>
        </Menu.Item>
        <Menu.Item key="4">
          <NavLink to="/signin">登录</NavLink>
        </Menu.Item>
          <Menu.Item key="5">
              <NavLink to="/signup">注册</NavLink>
          </Menu.Item>
        <Menu.Item key="6">
          <NavLink to="/account">我的</NavLink>
        </Menu.Item>
          <Menu.Item key="7">
              <Link to="/signout">退出</Link>
          </Menu.Item>
      </Menu>
    )
  }
}
