export const getItems = () => {
  const items: Item[] = [
    {
      id: 0,
      content: 'firstItem',
      filters: [{ id: 0, name: 'filterOne' }],
      createdAt: '',
      updatedAt: '',
    },
    {
      id: 1,
      content: 'secondItem',
      filters: [{ id: 1, name: 'filterTwo' }],
      createdAt: '',
      updatedAt: '',
    },
    {
      id: 2,
      content: 'thirdItem',
      filters: [{ id: 2, name: 'filterThree' }],
      createdAt: '',
      updatedAt: '',
    },
  ];

  return { items };
};
