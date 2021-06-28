import React from 'react'

import { AppContent } from '../../components/app-content'
import { Header } from '../../components/header'
import { HeroBanner } from '../../components/hero-banner'
import { Ranking } from '../../components/ranking'

export function HomePage() {
  return (
    <AppContent
      header={<Header />}
    >
      <HeroBanner />
      <Ranking />
    </AppContent>
  )
}
