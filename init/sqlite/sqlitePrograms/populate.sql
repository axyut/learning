.tables

insert into user ("id", "username", "password") values 
(1, "girish", "giriskK"),
(2, "gamvir", "gunamuna"),
(3, "samir", "123school"),
(4, "json", "hero@k"),
(5, "vabisya", "mahat999"),
(6, "sacar", "232www"),
(7, "vasme", "donHoMa"),
(8, "girish", "sherpa"),
(9, "sankhar", "hardikKK"),
(10, "sherpa", "gaunleKTO")
;

insert into post ("id", "user_id", "title", "description") values 
(1, 1, "SomeBody Hey", "Lorem ipsum is placeholder text commonly used in the graphic, print, and publishing industries for previewing layouts and visual mockups."),
(2,1, "gamvir", "Lorem ipsum is placeholder text commonly used in the graphic, print, and publishing industries for previewing layouts and visual mockups."),
(3,1, "samir", "Lorem ipsum is placeholder text commonly used in the graphic, print, and publishing industries for previewing layouts and visual mockups."),
(4,1, "json", "Lorem ipsum is placeholder text commonly used in the graphic, print, and publishing industries for previewing layouts and visual mockups."),
(5,1, "vabisya", "Lorem ipsum is placeholder text commonly used in the graphic, print, and publishing industries for previewing layouts and visual mockups."),
(6,1, "sacar", "Lorem ipsum is placeholder text commonly used in the graphic, print, and publishing industries for previewing layouts and visual mockups."),
(7,1, "vasme", "Lorem ipsum is placeholder text commonly used in the graphic, print, and publishing industries for previewing layouts and visual mockups."),
(8,1, "girish", "Lorem ipsum is placeholder text commonly used in the graphic, print, and publishing industries for previewing layouts and visual mockups."),
(9,1, "sankhar", "Lorem ipsum is placeholder text commonly used in the graphic, print, and publishing industries for previewing layouts and visual mockups."),
(10,1, "sherpa", "Lorem ipsum is placeholder text commonly used in the graphic, print, and publishing industries for previewing layouts and visual mockups."),
(11,2, "json", "Lorem ipsum is placeholder text commonly used in the graphic, print, and publishing industries for previewing layouts and visual mockups."),
(12,2, "vabisya", "Lorem ipsum is placeholder text commonly used in the graphic, print, and publishing industries for previewing layouts and visual mockups."),
(13,2, "sacar", "Lorem ipsum is placeholder text commonly used in the graphic, print, and publishing industries for previewing layouts and visual mockups."),
(14,2, "vasme", "Lorem ipsum is placeholder text commonly used in the graphic, print, and publishing industries for previewing layouts and visual mockups.")
;


insert into comment ("id", "post_id", "user_id", "title") values 
(1,1,1,"Mycomment"),
(2, 1,2, "My next comment"),
(3,1, 2, "Mycomment"),
(4,2,1, "Mycomment"),
(5,3,1, "Mycomment"),
(6,4,1, "Mycomment"),
(7,5,1,"Mycomment"),
(8,6,1, "Mycomment")
;

insert into follower ("user_id", "follower_id") values 
(1, 2),
(1,3),
(1,4),
(2,1),
(2,3),
(3,1),
(3,2),
(4,1)
;

insert into following ("user_id", "following_id") values 
(1, 3),
(1, 2),
(1, 4),
(2, 1),
(2, 3),
(2,4),
(3,1),
(3,2)
;

