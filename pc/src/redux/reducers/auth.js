import {
    SIGNIN_SUCCESS,
    SIGNIN_FAIL,
    SIGNOUT,
} from '../actions/types'

const auth = JSON.parse(localStorage.getItem("M-Token"));

const initAuth = auth ? auth : {
    token: {
        access_token: "",
        token_type: "",
        expires_in: 0,
    },
    // user: {
    //     username: ""
    // },
}
const auth = (state = initAuth, action) => {
    switch (action.type) {
        case SIGNIN_SUCCESS:
            console.log(action)
            return {
                ...state,
                token: {
                    access_token: action.payload.access_token,
                    token_type: action.payload.token_type,
                    expires_in: action.payload.expires_in,
                }
            }
        case SIGNIN_FAIL:
            return {...state, token: null, user: null}
        case SIGNOUT:
            return {...state, token: null, user: null}
        default:
            return state
    }
}

export default auth
