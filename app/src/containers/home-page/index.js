import React from 'react'
import {
  gql,
  useMutation,
  useQuery,
} from '@apollo/client'

import { AppContent } from '../../components/app-content'
import { Header } from '../../components/header'
import { HeroForm } from '../../components/hero-form'
import { RankingItem } from '../../components/ranking-item'

export function HomePage() {
  const { data, loading, error } = useQuery(HomePageQuery)

  if (loading) {
    return null
  }
  if (error) {
    return null
  }

  return (
    <AppContent
      header={<Header />}
    >
      <HeroForm />
      {data.rankings.map((r) => (
        <RankingItem key={r.username} { ...r } />
      ))}
    </AppContent>
  )
}

const HomePageQuery = gql`
  query HomePageQuery {
    rankings {
      ...RankingItemRankingFragment
    }
  }
  ${RankingItem.fragments.ranking}
`
