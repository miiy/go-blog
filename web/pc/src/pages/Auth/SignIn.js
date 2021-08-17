import React from 'react'
import {connect} from 'react-redux'
import { Form, Input, Button, Checkbox } from 'antd'
import { UserOutlined, LockOutlined } from '@ant-design/icons';

import {signIn} from '../../redux/actions/auth'


class SignIn extends React.Component {
    state = {
        username: "",
        password: ""
    }

    onFinish = (values) => {
        this.state.username = values.username
        this.state.password = values.password
        console.log('Received values of form: ', values);
        console.log(this.state)
    };

    signIn = () => {
        this.props.signIn(this.state)
    }
  render() {
    return (
        <Form
            name="normal_login"
            className="login-form"
            initialValues={{ remember: true }}
            onFinish={this.onFinish}
        >
            <Form.Item
                name="username"
                rules={[{ required: true, message: 'Please input your Username!' }]}
            >
                <Input prefix={<UserOutlined className="site-form-item-icon" />} placeholder="Username" />
            </Form.Item>
            <Form.Item
                name="password"
                rules={[{ required: true, message: 'Please input your Password!' }]}
            >
                <Input
                    prefix={<LockOutlined className="site-form-item-icon" />}
                    type="password"
                    placeholder="Password"
                />
            </Form.Item>
            <Form.Item>
                <Form.Item name="remember" valuePropName="checked" noStyle>
                    <Checkbox>Remember me</Checkbox>
                </Form.Item>

                <a className="login-form-forgot" href="">
                    Forgot password
                </a>
            </Form.Item>

            <Form.Item>
                <Button type="primary" htmlType="submit" className="login-form-button" onClick={this.signIn}>
                    Log in
                </Button>
                Or <a href="">register now!</a>
            </Form.Item>
        </Form>
    )
  }
}

export default connect(
    state => ({user: state.user}),
    {signIn}
)(SignIn)
