package me

// Add2qForm get all form
type Add2qForm struct {
	UserID uint `json:"-"`
	Twoq1  int  `json:"twoq1"`
	Twoq2  int  `json:"twoq2"`
	T2q    int  `json:"t2q"`
}

type Add9qForm struct {
	UserID   uint `json:"-"`
	Nineq1   int  `json:"nineq1"`
	Nineq2   int  `json:"nineq2"`
	Nineq3   int  `json:"nineq3"`
	Nineq4   int  `json:"nineq4"`
	Nineq5   int  `json:"nineq5"`
	Nineq6   int  `json:"nineq6"`
	Nineq7   int  `json:"nineq7"`
	Nineq8   int  `json:"nineq8"`
	Nineq9   int  `json:"nineq9"`
	T9q      int  `json:"t9q"`
	Status9q int  `json:"status9q"`
}

type Add8qForm struct {
	UserID   uint `json:"-"`
	Eightq1  int  `json:"eightq1"`
	Eightq2  int  `json:"eightq2"`
	Eightq3  int  `json:"eightq3"`
	Eightq3s  int  `json:"eightq3s"`
	Eightq4  int  `json:"eightq4"`
	Eightq5  int  `json:"eightq5"`
	Eightq6  int  `json:"eightq6"`
	Eightq7  int  `json:"eightq7"`
	Eightq8  int  `json:"eightq8"`
	T8q      int  `json:"t8q"`
	Status8q int  `json:"status8q"`
}
