import React from 'react';

export default function PlayerDetails({
  playerDisplayName,
  clockTime,
}: {
  playerDisplayName: string;
  clockTime: string;
}) {
  return (
    <div className="row">
      <p className="name">{playerDisplayName}</p>
      <p className="clock">{clockTime}</p>
    </div>
  );
}
