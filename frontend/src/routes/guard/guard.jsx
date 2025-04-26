import React from 'react';
import {Navigate, Outlet, useLocation} from 'react-router-dom';
import {SDB} from "../../shared/token.js";

function RouteGuard() {
    const location = useLocation();
    const token = SDB.get("token");

    if (!token) {
        // 重定向到登录页面，并在登录后返回之前访问的页面
        return <Navigate to="/login" state={{from: location}} replace/>;
    }

    // 如果用户已登录，渲染子路由（即ManageLayout）
    return <Outlet/>;
}

export default RouteGuard;
