import { HomePage } from './containers/home-page'
import { PlayPage } from './containers/play-page'
import { AuthenticationPage } from './containers/authentication-page'

import { AppLayout } from './components/app-layout'

export const routes = [
  {
    component: HomePage,
    exact: true,
    layout: AppLayout,
    path: '/',
  },
  {
    component: PlayPage,
    exact: true,
    layout: AppLayout,
    path: '/play',
  },
  {
    component: AuthenticationPage,
    exact: true,
    layout: AppLayout,
    path: '/a/:id/:token',
  },
]
