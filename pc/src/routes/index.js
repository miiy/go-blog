import {
  Home,
  About,
  SignIn,
  SignUp,
  NotFound,
  Articles,
  Article,
  Account,
  Feedback,
  Words,
  Posts
} from '../pages'

const routes = [
  {
    path: '/',
    exact: true,
    component: Home
  },
  {
    path: '/signin',
    component: SignIn
  },
  {
    path: '/signup',
    component: SignUp
  },
  {
    path: '/about',
    component: About
  },
  {
    path: '/404',
    component: NotFound
  },
  {
    path: '/articles',
    exact: true,
    component: Articles
  },
  {
    path: '/articles/:articleId',
    component: Article
  },
  {
    path: '/account',
    component: Account,
    routes: [
      {
        path: '/account/words',
        component: Words
      },
      {
        path: '/account/posts',
        component: Posts
      },
      {
        path: '/account/feedback',
        component: Feedback
      },
      {
        path: '*',
        component: NotFound
      },
    ]
  },
]

export default routes
