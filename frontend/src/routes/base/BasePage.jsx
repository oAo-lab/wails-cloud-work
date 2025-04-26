import React, {useMemo, useState} from 'react';
import {CloudFilled, MenuFoldOutlined, MenuUnfoldOutlined, RobotOutlined, SettingFilled} from '@ant-design/icons';
import {Button, Layout, Menu} from 'antd';

import '../../css/demo.css'
import TaskPage from "./TaskPage.jsx";
import CodeEditor from "./CodePage.jsx";
import ConfigPage from "./ConfigPage.jsx";
import Index from "./IndexPage.jsx";

const {Header, Sider, Content} = Layout;

const menuItems = [
    {key: '1', icon: <SettingFilled/>, label: '主页'},
    {key: '2', icon: <RobotOutlined/>, label: '任务集'},
    {key: '3', icon: <CloudFilled/>, label: '编辑器'},
    {key: '4', icon: <SettingFilled/>, label: '配置'},
];

const ManageLayout = () => {
    const [collapsed, setCollapsed] = useState(false);
    const [selectedMenuItem, setSelectedMenuItem] = useState('1');

    const renderContent = useMemo(() => {
        switch (selectedMenuItem) {
            case '1':
                return <Index/>
            case '2':
                return <TaskPage/>;
            case '3':
                return <CodeEditor/>;
            case '4':
                return <ConfigPage/>;
            default:
                return <div>请选择一个菜单项</div>;
        }
    }, [selectedMenuItem]);

    return (
        <Layout>
            <Sider trigger={null} collapsible collapsed={collapsed}>
                <Menu
                    theme="dark"
                    mode="vertical"
                    defaultSelectedKeys={['1']}
                    items={menuItems.map(item => ({
                        ...item,
                        onClick: () => setSelectedMenuItem(item.key),
                    }))}
                />
            </Sider>
            <Layout>
                <Header className='base-header-div' style={{padding: 0}}>
                    <Button
                        type="text"
                        icon={collapsed ? <MenuUnfoldOutlined/> : <MenuFoldOutlined/>}
                        onClick={() => setCollapsed(!collapsed)}
                        style={{fontSize: '16px', width: 32, height: 32}}
                    />
                    <img src='https://v2.jinrishici.com/one.svg' alt='loading...'/>
                </Header>
                <Content style={{margin: '24px 16px', minHeight: 280}}>
                    {renderContent}
                </Content>
            </Layout>
        </Layout>
    );
}

export default ManageLayout;
