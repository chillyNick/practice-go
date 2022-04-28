package chess

import "errors"

func CanKnightAttack(white, black string) (bool, error) {
	if err := arePositionValid(white, black); err != nil {
		return false, err
	}

	xd := absDiff(white[0], black[0])
	yd := absDiff(white[1], black[1])

	return xd == 2 && yd == 1 || xd == 1 && yd == 2, nil
}

func absDiff(x, y uint8) uint8 {
	if x > y {
		return x - y
	}
	return y - x
}

func arePositionValid(white, black string) error {
	if white == black {
		return errors.New("knights have to be placed in the different cells")
	}

	if err := isPositionValid(white); err != nil {
		return err
	}
	if err := isPositionValid(black); err != nil {
		return err
	}

	return nil
}

func isPositionValid(pos string) error {
	if len(pos) != 2 || !('a' <= pos[0] && pos[0] <= 'h') || !('1' <= pos[1] && pos[1] <= '8') {
		return errors.New("incorrect position")
	}

	return nil
}
