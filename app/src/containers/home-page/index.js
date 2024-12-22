import React from 'react'
import {
  gql,
  useMutation,
  useQuery,
} from '@apollo/client'

import { AppContent } from '../../components/app-content'
import { Footer } from '../../components/footer'
import { Header } from '../../components/header'
import { HeroForm } from '../../components/hero-form'
import { RankingItem } from '../../components/ranking-item'

export function HomePage() {
  const { data: query, loading, error } = useQuery(HomePageQuery)
  const [getStarted, { data: mutation }] = useMutation(HomePageGetStartedMutation)

  if (loading) {
    return null
  }
  if (error) {
    console.error(error)
    return null
  }

  return (
    <AppContent
      header={<Header />}
      footer={<Footer />}
    >
      <HeroForm { ...{ getStarted, ...mutation } } />
      {query.rankings.map((r) => (
        <RankingItem key={r.id} { ...r } />
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

const HomePageGetStartedMutation = gql`
  mutation HomePageGetStartedMutation(
    $email: String!
  ) {
    response: getStarted(email: $email)
  }
`
