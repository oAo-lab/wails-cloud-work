import React, {useState} from 'react';
import {Button, Form, Input, Modal} from 'antd';
import '../../css/demo.css'

const SectionFormModal = ({sectionName, visible, onCreate, onCancel}) => {
    const [form] = Form.useForm();

    // 动态渲染表单字段
    const renderFormFields = (section) => {
        switch (section) {
            case 'Email':
                return (
                    <>
                        <Form.Item name={['email', 'host']} label="主机" rules={[{required: true}]}>
                            <Input/>
                        </Form.Item>
                        <Form.Item name={['email', 'sender']} label="发送者" rules={[{required: true}]}>
                            <Input/>
                        </Form.Item>
                        <Form.Item name={['email', 'key']} label="秘钥" rules={[{required: true}]}>
                            <Input/>
                        </Form.Item>
                    </>
                );
            case 'HTML URL':
                return (
                    <>
                        <Form.Item name={['html_url', 'base_url']} label="就业平台地址" rules={[{required: true}]}>
                            <Input/>
                        </Form.Item>
                        <Form.Item name={['html_url', 'token_url']} label="验证码 Token 地址"
                                   rules={[{required: true}]}>
                            <Input/>
                        </Form.Item>
                        <Form.Item name={['html_url', 'code_url']} label="验证码识别地址" rules={[{required: true}]}>
                            <Input/>
                        </Form.Item>
                    </>
                );
            case 'Open-AI':
                return (
                    <>
                        <Form.Item name={['open_ai', 'url']} label="代理地址">
                            <Input
                                allowClear
                                defaultValue="https://api.aigcbest.top/"/>
                        </Form.Item>
                        <Form.Item name={['open_ai', 'key']} label="API Key" rules={[{required: true}]}>
                            <Input.Password/>
                        </Form.Item>
                    </>
                );
            case 'QQ-Bot':
                return (
                    <>
                        <Form.Item name={['qq_bot', 'appid']} label="AppID" rules={[{required: true}]}>
                            <Input.Password/>
                        </Form.Item>
                        <Form.Item name={['qq_bot', 'token']} label="Token" rules={[{required: true}]}>
                            <Input.Password/>
                        </Form.Item>
                    </>
                );
            default:
                return null;
        }
    };

    return (
        <Modal
            open={visible}
            style={{textAlign: "center"}}
            title={sectionName}
            okText="提交"
            cancelText="取消"
            onCancel={onCancel}
            onOk={() => {
                form.validateFields()
                    .then((values) => {
                        form.resetFields();
                        onCreate(values);
                    })
                    .catch((info) => {
                        console.log('Validate Failed:', info);
                    });
            }}
        >
            <Form
                form={form}
                layout="vertical"
                name={`form_in_modal_${sectionName}`}
            >
                {renderFormFields(sectionName)}
            </Form>
        </Modal>
    );
};

const ConfigPage = () => {
    const [visible, setVisible] = useState(false);
    const [currentSection, setCurrentSection] = useState('');

    const onCreate = (values) => {
        console.log(`${currentSection} Received values:`, values);
        setVisible(false);
    };

    const openModal = (sectionName) => {
        setCurrentSection(sectionName);
        setVisible(true);
    };

    return (
        <div className='config-container'>
            <Button onClick={() => openModal('Email')}>邮箱</Button>
            <Button onClick={() => openModal('HTML URL')}>页面</Button>
            <Button onClick={() => openModal('Open-AI')}>Open-AI</Button>
            <Button onClick={() => openModal('QQ-Bot')}>QQ-Bot</Button>


            <SectionFormModal
                sectionName={currentSection}
                visible={visible}
                onCreate={onCreate}
                onCancel={() => setVisible(false)}
            />
        </div>
    );
};

export default ConfigPage;
