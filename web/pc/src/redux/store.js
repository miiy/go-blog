import {createStore, applyMiddleware} from 'redux'
// https://github.com/zalmoxisus/redux-devtools-extension#installation
import { composeWithDevTools } from 'redux-devtools-extension';
import thunk from 'redux-thunk'

import reducers from './reducers'

export default createStore(reducers, composeWithDevTools(applyMiddleware(thunk)))