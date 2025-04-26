import React, {useState} from 'react'
import {Button, Card, Form, Input, message} from "antd";

import "../../css/demo.css"
import {LockOutlined, UserOutlined} from "@ant-design/icons";
import request from "../../shared/request.js";
import {SDB} from "../../shared/token.js";
import {useNavigate} from "react-router-dom";


const LoginPage = () => {

    // 页面跳转
    const navigation = useNavigate()
    // 按钮状态
    const [clickedButton, setClickedButton] = useState(null)

    const handleLogin = () => setClickedButton(true)
    const handleSignup = () => setClickedButton(false)


    const handleFormSubmit = async (values) => {
        let resp = {msg: ''}

        if (clickedButton) {
            resp = await request.loginUser(values)

            if (resp.code === 200) {
                SDB.set("token", resp.data.token);
                setTimeout(() => navigation('/'), 0)
                message.success(resp.msg); // 移除await
            } else {
                message.error(resp.msg);
            }
        } else {
            resp = await request.registerUser(values)
            if (resp.code === 200) {
                message.success(resp.msg)
                SDB.set("user_id", resp.data.token)
            } else {
                message.error(resp.msg)
            }
        }
    }

    return (
        <>
            <div className='login-div'>
                <Card className='login-card'>
                    <FormLayout
                        Finish={handleFormSubmit}
                        Login={handleLogin}
                        Signup={handleSignup}
                    />
                </Card>
            </div>
        </>
    )
}

let FormLayout = ({Finish, Login, Signup}) => {
    return (
        <Form
            name="normal_login"
            className="login-form"
            initialValues={{remember: true}}
            onFinish={Finish}
        >
            <Form.Item style={{textAlign: 'center'}}>
                <strong style={{fontSize: '20px'}}>
                    登录
                </strong>
            </Form.Item>

            <Form.Item name="username" rules={[{required: true, message: '请输入账号'}]}>
                <Input
                    prefix={<UserOutlined className="site-form-item-icon"/>}
                    placeholder="Username"
                    allowClear="true"
                />
            </Form.Item>

            <Form.Item name="password" rules={[{required: true, message: '请输入密码'}]}>
                <Input
                    prefix={<LockOutlined className="site-form-item-icon"/>}
                    type="password"
                    placeholder="admin"
                />
            </Form.Item>

            <Form.Item>
                <div style={{display: 'flex', justifyContent: 'space-between'}}>
                    <CustomButton
                        buttonText="登录"
                        buttonType="primary"
                        onClick={Login} // 正确：传递处理函数
                        showFlag="true"
                    />
                    <CustomButton
                        buttonText="注册"
                        buttonType="default"
                        onClick={Signup} // 正确：传递处理函数
                    />
                </div>
            </Form.Item>
        </Form>
    )
}


let CustomButton = ({buttonText, buttonType, onClick, showFlag}) => {
    const [isHovered, setIsHovered] = useState(false)

    return (
        <Button
            block
            shape="round"
            type={buttonType}
            htmlType="submit"
            style={{flex: 1, marginRight: '8px',}}
            onMouseEnter={() => setIsHovered(true)}
            onMouseLeave={() => setIsHovered(false)}
            onClick={onClick}
        >
            {buttonText}
            {isHovered && showFlag ? '🥕' : ''}
        </Button>
    )
}


export default LoginPage