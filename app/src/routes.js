import { HomePage } from './containers/home-page'
import { PlayPage } from './containers/play-page'

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
]
