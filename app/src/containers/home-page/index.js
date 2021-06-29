import React from 'react'

import { AppContent } from '../../components/app-content'
import { Header } from '../../components/header'
import { HeroForm } from '../../components/hero-form'
import { Ranking } from '../../components/ranking'

export function HomePage() {
  return (
    <AppContent
      header={<Header />}
    >
      <HeroForm />
      <Ranking />
    </AppContent>
  )
}
