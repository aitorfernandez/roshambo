import { HomePage } from './containers/home-page'

import { AppLayout } from './components/app-layout'

export const routes = [
  {
    component: HomePage,
    exact: true,
    layout: AppLayout,
    path: '/',
  },
]
