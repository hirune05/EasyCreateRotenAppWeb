

type LoginStudentResponse = {
  Token:      string;
  StudentID:  int;
  Name:       string;
  StoreID:    int;
  StoreName:  string;
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

type Student = {
  ID:        int;
  Name:      string;
  Password:  string;
  CreatedAt: Date;
};

