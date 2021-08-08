import React from 'react';

import { SmileOutlined, CrownOutlined, TabletOutlined, AntDesignOutlined } from '@ant-design/icons';

export default {
    title: 'ESP',
    route: {
        path: '/',
        routes: [
            {
                path: '/welcome',
                name: '欢迎',
                icon: <SmileOutlined />,
                component: './Welcome',
            },
            {
                name: '公众服务',
                icon: <TabletOutlined />,
                path: '/inclusive',
                component: './ListTableList',
                routes: [
                    {
                        path: '/inclusive/shenghuo',
                        name: '生活服务',
                        icon: <CrownOutlined />,
                        routes: [
                            {
                                path: 'lifa',
                                name: '理发',
                                icon: <CrownOutlined />,
                                component: './Welcome',
                            },
                            {
                                path: 'baojie',
                                name: '保洁',
                                icon: <CrownOutlined />,
                                component: './Welcome',
                            },
                            {
                                path: 'songcan',
                                name: '送餐',
                                icon: <CrownOutlined />,
                                component: './Welcome',
                            },
                        ],
                    },
                    {
                        path: '/inclusive/university',
                        name: '老年大学',
                        icon: <CrownOutlined />,
                        component: './Welcome',
                        routes: [
                            {
                                path: 'register',
                                name: '报名',
                                icon: <CrownOutlined />,
                                component: './Welcome',
                            },
                            {
                                path: 'study',
                                name: '学习',
                                icon: <CrownOutlined />,
                                component: './Welcome',
                            },
                            {
                                path: 'music',
                                name: '音乐',
                                icon: <CrownOutlined />,
                                component: './Welcome',
                            },
                            {
                                path: 'handwrite',
                                name: '书法',
                                icon: <CrownOutlined />,
                                component: './Welcome',
                            },
                            {
                                path: 'panting',
                                name: '绘画',
                                icon: <CrownOutlined />,
                                component: './Welcome',
                            },
                        ],
                    },
                    {
                        path: '/inclusive/health',
                        name: '健康管理',
                        icon: <CrownOutlined />,
                        component: './Welcome',
                        routes: [
                            {
                                path: 'chronic',
                                name: '慢病讲座',
                                icon: <CrownOutlined />,
                                component: './Welcome',
                            },
                            {
                                path: 'tourguide',
                                name: '健康出游',
                                icon: <CrownOutlined />,
                                component: './Welcome',
                            },
                        ],
                    },
                ]
            },
            {
                path: '/premium',
                name: '会员服务',
                icon: <CrownOutlined />,
                access: 'canAdmin',
                component: './premium',
                routes: [
                    {
                        path: '/premium/health',
                        name: '健康管理',
                        icon: <CrownOutlined />,
                        component: './Welcome',
                        routes:[
                            {
                                path: 'tizheng',
                                name: '体征监测',
                                icon: <CrownOutlined />,
                                component: './Welcome',
                            },
                            {
                                path: 'shenghuo',
                                name: '生活监测(非视频)',
                                icon: <CrownOutlined />,
                                component: './Welcome',
                            },
                        ]
                    },
                    {
                        path: '/premium/laoyou',
                        name: '老幼同营',
                        icon: <CrownOutlined />,
                        component: './Welcome',
                    },
                    {
                        path: '/premium/yiliao',
                        name: '医疗优先',
                        icon: <CrownOutlined />,
                        component: './Welcome',
                        routes:[
                            {
                                path: 'volunteer',
                                name: '志愿者服务',
                                icon: <CrownOutlined />,
                                component: './Welcome',
                            },
                            {
                                path: 'chuxing',
                                name: '出行陪伴',
                                icon: <CrownOutlined />,
                                component: './Welcome',
                            },
                        ]
                    },
                ],
            },
            {
                path: '/admin',
                name: '管理',
                icon: <AntDesignOutlined />,
                access: 'canAdmin',
                component: './Admin',
                routes: [
                    {
                        path: '/admin/info',
                        name: '个人信息',
                        icon: <AntDesignOutlined />,
                        component: './Welcome',
                    },
                    {
                        path: '/admin/help',
                        name: '使用帮助',
                        icon: <AntDesignOutlined />,
                        component: './Welcome',
                    },
                    {
                        path: '/admin/about',
                        name: '关于',
                        icon: <AntDesignOutlined />,
                        component: './Welcome',
                    },
                ],
            },

        ]
    },
    location: {
        pathname: '/',
    },
};