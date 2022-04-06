import logo from './logo.svg';
import './App.less';

// store
import {Provider} from 'react-redux'
import store from './redux/store'

// router
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link
} from 'react-router-dom'
import routes from './routes'

import Nav from './components/Nav/Nav'

import { Layout, Menu, Breadcrumb } from 'antd';
const { Header, Content, Footer } = Layout;

function App() {
    return (
        <Provider store={store}>
            <Router>
                <Layout>
                    <Header style={{ position: 'fixed', zIndex: 1, width: '100%' }}>
                        {/*<div className="logo">logo</div>*/}
                        <Nav />
                    </Header>
                    <Content className="site-layout" style={{ padding: '0 50px', marginTop: 64 }}>
                        <Breadcrumb style={{ margin: '16px 0' }}>
                            <Breadcrumb.Item>Home</Breadcrumb.Item>
                            <Breadcrumb.Item>List</Breadcrumb.Item>
                            <Breadcrumb.Item>App</Breadcrumb.Item>
                        </Breadcrumb>
                        <Switch>
                            {
                                routes.map((route, i) => {
                                    return <Route key={i} exact={route.exact} path={route.path} component={route.component} />
                                })
                            }
                        </Switch>
                    </Content>
                    <Footer style={{ textAlign: 'center' }}>Footer</Footer>
                </Layout>
            </Router>
        </Provider>
    )
}

export default App;
