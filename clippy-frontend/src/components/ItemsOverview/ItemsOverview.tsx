import React from 'react';
import { useQuery } from '@tanstack/react-query';
import { getItems } from '../../services/ItemsService';

export const ItemsOverview = () => {
  const { data } = useQuery(['items'], getItems);

  return (
    <ul>
      {data?.items.map(({ id, content }) => (
        <li key={id}>{content}</li>
      ))}
    </ul>
  );
};
