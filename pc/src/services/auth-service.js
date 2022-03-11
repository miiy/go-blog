import axios from "../axios/axios";

class AuthService {
    signUp = (signUpForm) => {
        return axios.post('/api/v1/auth/signup', signUpForm)
    }

    signIn = ({username, password}) => {
        return axios.post('/api/v1/auth/signin', {username, password}).then((response) => {
            if (response.data.accessToken) {
                localStorage.setItem("user", JSON.stringify(response.data));
            }
            return response.data
        });
    }

    signOut = () => {
        localStorage.removeItem("user")
    }

    authInfo = () => {
        return axios.post('/api/v1/auth/me')
    }
}

export default new AuthService();