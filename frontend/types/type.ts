


type LoginStudentResponse = {
  Token:      string;
  StudentID:  int;
  Name:       string;
};

type AdminUser = {
  ID: int:
  Username: string;
  Password: string;
  Email:    string;
  CreatedAt: Date;
  UpdatedAt: Date;
};


type Event = {
  ID:        int;
  Name:      string;
  Year:      int;
  StartTime: Date;
  EndTime:   Date;
  CreatedAt: Date;
  UpdatedAt: Date;
};
