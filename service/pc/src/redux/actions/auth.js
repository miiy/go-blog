import AuthService from '../../services/auth-service'

import {SIGNIN_SUCCESS, SIGNIN_FAIL} from './types'

export const authSuccess = (tokenInfo) => {
    return {
        type: SIGNIN_SUCCESS,
        payload: tokenInfo,
    }
}

export const authError = (msg) => {
    return {
        type: SIGNIN_FAIL,
        msg,
    }
}

export const signIn = ({username, password}) => (dispatch) => {
    return AuthService.signIn({username, password}).then((resp) => {
            localStorage.setItem('M-Token', JSON.stringify(resp))
            dispatch(authSuccess(resp))
        return Promise.resolve(resp)
    })
}

export const authInfo = (token) => {
    return AuthService.authInfo(token)
}

export const signUp = (signUpForm) => (dispatch) => {
    return AuthService.signUp(signUpForm)
}

export const signOut = () => {
    localStorage.removeItem('M-Token')
}