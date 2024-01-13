import React from 'react'
import ReactDOM from 'react-dom/client'
import {createBrowserRouter, RouterProvider} from 'react-router-dom'

import './css/index.css'
import RouteGuard from "./routes/guard/guard.jsx"
import ManageLayout from './routes/base/BasePage'
import LoginPage from "./routes/base/LoginPage.jsx";
import NotFoundPage from "./routes/guard/404.jsx";

const router = createBrowserRouter([
    {
        path: "/",
        element: <RouteGuard/>,
        children: [
            {
                path: "/",
                element: <ManageLayout/>,
            },
        ]
    },
    {
        path: "/login",
        element: <LoginPage/>,
    },
    {
        path: "*",
        element: <NotFoundPage/>,
    },
]);

ReactDOM.createRoot(document.getElementById('root')).render(
    <React.StrictMode>
        <RouterProvider router={router}/>
    </React.StrictMode>
)
