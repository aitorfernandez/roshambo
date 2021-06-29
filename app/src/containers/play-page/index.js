import React from 'react'

import { AppContent } from '../../components/app-content'
import { Header } from '../../components/header'
import { Profile } from '../../components/profile'
import { StatCreateButton } from '../../components/stat-create-button'
import { StatDeleteButton } from '../../components/stat-delete-button'
import { StatItem } from '../../components/stat-item'
import { StatResult } from '../../components/stat-result'
import { StatScore } from '../../components/stat-score'

export function PlayPage() {
  return (
    <AppContent
      header={
        <>
          <Header />
          <Profile />
        </>
      }
    >
      <StatResult />
      <StatScore />
      <div>
        <StatCreateButton />
        <StatCreateButton />
        <StatCreateButton />
        <StatDeleteButton />
      </div>
      <StatItem />
    </AppContent>
  )
}
