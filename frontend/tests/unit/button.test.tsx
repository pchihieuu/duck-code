import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/react';
import { Button } from '@/components/ui/button';

describe('Button', () => {
  it('renders children', () => {
    render(<Button>Nhấn vào đây</Button>);
    expect(screen.getByText('Nhấn vào đây')).toBeInTheDocument();
  });
});
