import React from 'react'
import {Button, notification} from "antd"
import request from "../../shared/request.js"
import {SDB} from "../../shared/token.js";
import '../../css/demo.css'
import {useNavigate} from "react-router-dom";

export default function Index() {
    const navigation = useNavigate()
    const [api, contextHolder] = notification.useNotification();

    const openNotify = (title, desc) => {
        api.open({
            message: title,
            description: desc,
            duration: 2.5,
        });
    };

    return (
        <div>
            <div className={'index-container'}>
                <h3>
                    <strong>功能测试</strong>
                </h3>
                <span>
                Github地址: <a href="https://github.com/Fromsko">FromSko</a>
                </span>
            </div>

            <div className={'config-container'}>
                <CustomButton text='注册' onClick={async () => {
                    openNotify("dd", "ddd")
                    await authRegister()
                }}/>
                <CustomButton text='登录' onClick={() => authLogin()}/>
                <CustomButton text='更新' onClick={() => authUpdate()}/>
                <CustomButton text='删除' onClick={() => authDelete()}/>
                <CustomButton text='刷新' onClick={() => {
                    SDB.remove('token')
                    navigation("/login")
                }}/>
            </div>
        </div>
    )
}

let CustomButton = ({text, onClick}) => {
    return (
        <Button type='default' onClick={onClick}>
            {text}
        </Button>
    )
}

const authRegister = async () => {
    let userInfo = {
        username: 'admin',
        password: 'admin'
    }
    let resp = await request.registerUser(userInfo)
    console.log(resp)
}

const authLogin = async () => {
    let userInfo = {
        username: 'admin',
        password: 'admin'
    }
    let resp = await request.loginUser(userInfo)
    console.log(resp)
}

const authUpdate = async () => {
    let userInfo = {
        username: 'admin',
        password: 'admin'
    }
    let resp = await request.updateUser(SDB.get('user_id'), userInfo)
    console.log(resp)
}

const authDelete = async () => {
    let resp = await request.deleteUser(SDB.get('user_id'))
    console.log(resp)
}